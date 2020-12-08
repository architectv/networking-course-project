function setCookie(cname, cvalue, exsec) {
  var d = new Date();
  d.setTime(d.getTime() + (exsec*1000));
  var expires = "expires="+ d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getCookie(cname) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

export function validate_prop(validators, prop, value) {
  console.log("Validate", prop, value)
  if (validators[prop]) {
    return validators[prop](value);
  }
  return null;
}

export function validate(validators, data) {
  let res = {};
  let flag = false;
  for (const key in data) {
    let ret = validate_prop(key, data[key]);
    if (ret) {
      res[key] = ret;
      flag = true;
    }
  }
  if (flag) {
    return res;
  }
  return null;
}

export function getValidData(fields, obj) {
  let data = {}
  for (let i in fields) {
    data[fields[i].key] = fields[i].value;
  }
  let validation = obj.validate(data);
  if (validation) {
    return null;
  }
  return data;
}

export function validateField(obj, field, e) {
  console.log(field);
  if (e.srcElement) {
    let value = e.srcElement.value;
    let invalid = obj.validate_prop(field.key, value);
    if (invalid) {
      field.invalid = true;
      field.error = invalid;
      field = field;
    } else {
      field.invalid = false;
    }
  }
}

