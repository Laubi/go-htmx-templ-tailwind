/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/views/**/*.{go,js,templ,html}"
    ],
    theme: {
        extend: {},
    },
    plugins: [
        // adds CSS variables for colors
        // https://gist.github.com/Merott/d2a19b32db07565e94f10d13d11a8574?permalink_comment_id=4194451#gistcomment-4194451
        function ({ addBase, theme }) {
            function extractColorVars (colorObj, colorGroup = '') {
                return Object.keys(colorObj).reduce((vars, colorKey) => {
                    const value = colorObj[colorKey];
                    const cssVariable = colorKey === "DEFAULT" ? `--color${colorGroup}` : `--color${colorGroup}-${colorKey}`;

                    const newVars =
                        typeof value === 'string'
                            ? { [cssVariable]: value }
                            : extractColorVars(value, `-${colorKey}`);

                    return { ...vars, ...newVars };
                }, {});
            }

            addBase({
                ':root': extractColorVars(theme('colors')),
            });
        }
    ],
}