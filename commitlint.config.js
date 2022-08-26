// Note: Github actions use this config file to lint commit messages
const { rules } = require('@commitlint/config-conventional');

module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [2, 'always', [...rules['type-enum'][2], 'release']],
    'scope-enum': [2, 'always', ['core', 'client', 'gin-sample', 'deps']],
  },
};
