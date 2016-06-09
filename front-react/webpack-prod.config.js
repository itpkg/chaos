module.exports = require("./webpack.config")({
    minify: true,
    api: '/api/v1',
    engines: ['platform', 'cms', 'hr', 'ops', 'reading', 'team'],
    version: '0.0.1'
});
