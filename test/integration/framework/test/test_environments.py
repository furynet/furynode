import contextlib
from furytool.common import *
from furytool import command, environments, project, furynet, cosmos


def get_validators(env):
    furynoded = env._furynoded_for(env.node_info[0])
    return {v["description"]["moniker"]: v for v in furynoded.query_staking_validators()}


def test_transfer(env):
    furynoded = env.furynoded
    furynoded.send_and_check(env.faucet, furynoded.create_addr(), {furynet.FURY: 10 ** furynet.FURY_DECIMALS})


def assert_validators_working(env, expected_monikers):
    assert set(get_validators(env)) == expected_monikers
    for i in range(len(env.node_info)):
        test_transfer(env)


class TestFurynodedEnvironment:
    def setup_method(self):
        self.cmd = command.Command()
        self.furynoded_home_root = self.cmd.tmpdir("furytool.tmp")
        self.cmd.rmdir(self.furynoded_home_root)
        self.cmd.mkdir(self.furynoded_home_root)
        prj = project.Project(self.cmd, project_dir())
        prj.pkill()

    def teardown_method(self):
        prj = project.Project(self.cmd, project_dir())
        prj.pkill()

    def test_environment_setup_basic(self):
        env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
        env.add_validator()
        env.start()
        assert_validators_working(env, set("furynoded-{}".format(i) for i in range(1)))

    def test_add_validator_before_and_after_start(self):
        env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
        env.add_validator()
        env.add_validator()
        env.init()
        env.start()
        env.add_validator()
        assert_validators_working(env, set("furynoded-{}".format(i) for i in range(3)))

    def test_environment_fails_to_start_if_commission_rate_is_over_max(self):
        env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
        env.add_validator(commission_rate=0.10, commission_max_rate=0.05)
        exception = None
        try:
            env.start()
        except Exception as e:
            exception = e
        # The validator will exit immediately, writing error to the log.
        # What we get here is a "timeout waiting for furynoded to come up".
        assert type(exception) == furynet.FurynodedException

    def test_need_2_out_of_3_validators_running_for_consensus(self):
        env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
        env.add_validator()
        env.add_validator()
        env.add_validator()
        env.add_validator()
        env.start()
        assert len(env.running_processes) == 4
        test_transfer(env)  # 4 out of 4 => should work
        env.running_processes[-1].terminate()
        env.running_processes[-1].wait()
        env.open_log_files[-1].close()
        env.running_processes.pop()
        env.open_log_files.pop()
        test_transfer(env)  # 3 out of 4 => should work
        env.running_processes[-1].terminate()
        env.running_processes[-1].wait()
        env.open_log_files[-1].close()
        env.running_processes.pop()
        env.open_log_files.pop()
        exception = None
        try:
            test_transfer(env)  # 2 out of 4 => should fail
        except Exception as e:
            exception = e
        assert type(exception) == furynet.FurynodedException

    def test_can_have_validators_with_same_moniker(self):
        env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
        env.add_validator()
        env.start()
        furynoded = env.furynoded
        home1 = self.cmd.mktempdir()
        home2 = self.cmd.mktempdir()
        try:
            env.add_validator(moniker="juniper", home=home1)
            assert len(furynoded.query_staking_validators()) == 2
            env.add_validator(moniker="juniper", home=home2)
            assert len(furynoded.query_staking_validators()) == 3
        finally:
            self.cmd.rmdir(home1)
            self.cmd.rmdir(home2)

    def test_balances(self):
        number_of_validators = 3
        number_of_denoms = 10  # > 1
        number_of_wallets = 100
        faucet_balance = cosmos.balance_add({"foo{}".format(i): (i + 1) * 10**30 for i in range(10)},
            {furynet.FURY: 10**30})

        tmpdir = self.cmd.mktempdir()
        try:
            furynoded = furynet.Furynoded(self.cmd, home=tmpdir)
            extra_accounts = {furynoded.create_addr(): {"bar{}".format(j): (i * number_of_denoms + j + 1) * 10**25
                for j in range(number_of_denoms)} for i in range(number_of_wallets)}
            env = environments.FurynodedEnvironment(self.cmd, furynoded_home_root=self.furynoded_home_root)
            for _ in range(number_of_validators):
                env.add_validator()
            env.init(faucet_balance=faucet_balance, extra_accounts=extra_accounts)
            env.start()

            # Check faucet balances
            for i in range(number_of_validators):
                furynoded = env._furynoded_for(env.node_info[i])
                assert cosmos.balance_equal(furynoded.get_balance(env.faucet), faucet_balance)

            # Check balances of extra_accounts
            for i in range(number_of_validators):
                furynoded = env._furynoded_for(env.node_info[i])
                assert cosmos.balance_equal(furynoded.get_balance(env.faucet), faucet_balance)
                for addr, balance in extra_accounts.items():
                    assert cosmos.balance_equal(furynoded.get_balance(addr), balance)

            # Check funding
            for i in range(number_of_validators):
                addr = furynoded.create_addr()
                amount = {"foo0": 100 * 10**18, "foo1": 100 * 10**18}
                env.fund(addr, amount)
                assert cosmos.balance_equal(furynoded.get_balance(addr), amount)

            # On each node, do a sample transfer of one fury from admin to a new wallet and check that the change of
            # balances is visible on all nodes
            test_transfer_amount = {furynet.FURY: 10**furynet.FURY_DECIMALS}
            for i in range(number_of_validators):
                furynoded_i = env._furynoded_for(env.node_info[i])
                test_addr = furynoded_i.create_addr()
                # Sending also requires fury, this way we're making sure that each validator has some
                admin_addr = env.node_info[i]["admin_addr"]
                furynoded_i.send(admin_addr, test_addr, test_transfer_amount)
                for j in range(number_of_validators):
                    furynoded_j = env._furynoded_for(env.node_info[j])
                    furynoded_j.wait_for_balance_change(test_addr, {}, test_transfer_amount)
        finally:
            self.cmd.rmdir(tmpdir)
