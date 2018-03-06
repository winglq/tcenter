var path = require('path');
var output_dir = path.resolve("../../../resources/frontend/littlefriends/static/js/")
module.exports = {
    entry: './common/main.js',
    output: {
        path: output_dir,
        filename: 'common.js'
    },
    module: {
        rules: [
            { test: path.join(__dirname, 'common'),
              exclude: /node_modules/,
              loader: 'babel-loader' }
        ]
    }
};
