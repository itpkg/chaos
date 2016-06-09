module.exports = require("./webpack.config")({
    minify: true,
    backend: '/api/v1',
    engines: ['platform', 'cms', 'hr', 'ops', 'reading', 'team']
});
