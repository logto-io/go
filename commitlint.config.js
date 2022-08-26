const { rules } = require('@commitlint/config-conventional');

/** @type {import('@commitlint/types').UserConfig} **/
module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [2, 'always', [...rules['type-enum'][2], 'release']],
    'scope-enum': [2, 'always', ['core', 'client', 'gin-sample', 'deps']],
  },
};
