/* application.js -- the main js file*/ 

// jQuery (for Kube)
// require("expose-loader?$!expose-loader?jQuery!jquery");



const IPFS = require("ipfs");
const node = new IPFS()
console.log(node)

// require("bootstrap-sass/assets/javascripts/bootstrap.js");

// Libraries and Any JS (Requires/Imports)

$(() => {
  console.log('we outchea');

  // const heatmap_data = {
  //   "1485684800": 4,
  //   "1484684800": 4,
  //   "1484684800": 3,
  //   "1484684800": 4,
  //   "1484684800": 2,
  //   "1484684800": 2,
  //   "1484684800": 2,
  //   "1484684800": 1,
  //   "1484684800": 2,
  //   "1484684800": 5,
  //   "1486684800": 4,
  //   "1486771200": 2,
  //   "1489857600": 1,
  // }

  // let heatmap = new Chart({
  //   parent: "#gh-chart",
  //   type: "heatmap",
  //   height: 115,
  //   data: heatmap_data, // object with date/timestamp-value pairs

  //   // discrete_domains: 1, // default: 0

  //   // start: 1486684800,
  //   // A Date object;
  //   // default: today's date in past year
  //   //          for an annual heatmap

  //   legend_colors: ["#ebedf0", "#fdf436", "#ffc700", "#ff9100", "#06001c"]
  //   // Set of five incremental colors,
  //   // beginning with a low-saturation color for zero data;
  //   // default: ['#ebedf0', '#c6e48b', '#7bc96f', '#239a3b', '#196127']
  // });

}); // end of js