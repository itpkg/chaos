require("bootstrap/dist/css/bootstrap.css");
require("bootstrap/dist/css/bootstrap-theme.css");
require("./main.css");

import $ from 'jquery';
import React from 'react';

console.log("jquery version: "+$().jquery);
console.log("react version: "+React.version);
console.log("chaos version: "+CHAOS_ENV.version);
