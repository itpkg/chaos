module.exports = require("./webpack.config")({
    minify: true,
    host: '/api/v1',
    engines: ['platform', 'cms', 'hr', 'ops', 'reading', 'team']
});
