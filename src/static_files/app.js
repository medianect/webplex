/*
  Should maybe use addEventListener('onhover', show, false), this would help
  with uglifyjs and permit to not use the eslint disable. Maybe even to use
  sourceType: 'module' in .eslintrc.yaml
*/

function show (id) { document.getElementById(id).style.display = 'block' } // eslint-disable-line no-unused-vars
function hide (id) { document.getElementById(id).style.display = 'none' } // eslint-disable-line no-unused-vars
