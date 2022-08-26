# Install commitlint/cli to lint commit message
npm install -g @commitlint/config-conventional @commitlint/cli

# Link @commitlint/config-conventional to project
npm link @commitlint/config-conventional

# Install husky to add git hooks
npm install -g husky
husky install
