module.exports = {
  root: true,
  env: {
    node: true,
  },
  parserOptions: {
    parser: "@babel/eslint-parser",
  },
  extends: ["plugin:vue/recommended"],
  rules: {
    // Default JS
    indent: ["error", 2],
    semi: 'off',
    "space-before-function-paren": ["error", "never"],
    // Vue Plugin
    "vue/html-indent": ["error", 2],
    "vue/max-attributes-per-line": [2],
    "vue/component-name-in-template-casing": ["error", "kebab-case"],
    "vue/script-indent": ["error", 2, {baseIndent: 1}],
    "vue/singleline-html-element-content-newline": "off",
    "vue/multiline-html-element-content-newline": "off",
    "vue/multi-word-component-names": "off",
    "vue/component-tags-order": ["error", {
      "order": ["template", "script", "style"]
    }],
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'max-len': 'off',
    'linebreak-style': 'off',
    camelcase: ['error', {properties: 'never', ignoreDestructuring: true, ignoreImports: true}],
    'arrow-parens': ['error', 'as-needed'],
    'import/no-anonymous-default-export': 'off',
    'no-restricted-globals': 'off'
  },
  overrides: [
    {
      files: ["*.vue"],
      rules: {
        indent: "off",
      },
    },
  ],
}
