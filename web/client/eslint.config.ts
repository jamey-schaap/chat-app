import js from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";

export default tseslint.config([
    {
        files: ["**/*.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"], plugins: {js}, extends: [{
            name: "js/recommended"
        }], languageOptions: {globals: globals.browser}
    },
    tseslint.configs.recommended,
    pluginReact.configs.flat.recommended,
    {
        rules: {
            "react/react-in-jsx-scope": "off"
        }
    }
]);
