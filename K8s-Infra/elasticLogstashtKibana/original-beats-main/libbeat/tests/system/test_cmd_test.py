import os
import logging
import unittest
import pytest
from base import BaseTest


INTEGRATION_TESTS = os.environ.get('INTEGRATION_TESTS', False)


class TestCommandTest(BaseTest):
    """
    Test beat subcommands
    """

    def test_config(self):
        """
        Test test config command
        """
        self.render_config_template("mockbeat",
                                    os.path.join(self.working_dir, "libbeat.yml"))

        exit_code = self.run_beat(
            logging_args=[],
            extra_args=["test", "config"],
            config="libbeat.yml")

        assert exit_code == 0
        assert self.log_contains("Config OK")

    def test_bad_config(self):
        """
        Test test config command with bad config
        """
        exit_code = self.run_beat(
            logging_args=[],
            extra_args=["test", "config"],
            config="libbeat-missing.yml")

        assert exit_code == 1
        assert self.log_contains("Config OK") is False

    @unittest.skipUnless(INTEGRATION_TESTS, "integration test")
    @pytest.mark.tag('integration')
    def test_output(self):
        """
        Test test output works
        """

        self.es_client()
        logging.getLogger("elasticsearch").setLevel(logging.ERROR)

        self.render_config_template("mockbeat",
                                    os.path.join(self.working_dir, "mockbeat.yml"),
                                    elasticsearch=self.get_elasticsearch_template_config())
        exit_code = self.run_beat(
            extra_args=["test", "output"],
            config="mockbeat.yml")

        assert exit_code == 0
        assert self.log_contains('parse url... OK')
        assert self.log_contains('TLS... WARN secure connection disabled')
        assert self.log_contains('talk to server... OK')

    @unittest.skipIf(not INTEGRATION_TESTS, "integration test")
    def test_wrong_output(self):
        """
        Test test wrong output works
        """
        self.render_config_template("mockbeat",
                                    os.path.join(self.working_dir,
                                                 "mockbeat.yml"),
                                    elasticsearch={
                                        "host": 'badhost:9200',
                                        "user": 'admin',
                                        "pass": 'testing'
                                    })
        exit_code = self.run_beat(
            extra_args=["test", "output"],
            config="mockbeat.yml")

        assert exit_code == 1
        assert self.log_contains('parse url... OK')
        assert self.log_contains('dns lookup... ERROR')
