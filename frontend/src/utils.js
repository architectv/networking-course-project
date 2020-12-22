function setCookie(cname, cvalue, exsec) {
  var d = new Date();
  d.setTime(d.getTime() + (exsec*1000));
  var expires = "expires="+ d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

export function getDate(ts) {
  let date = new Date(ts * 1000);
  let day = "0" + date.getDate();
  let month = "0" + (date.getMonth() + 1);
  let year = date.getFullYear();
  let hours = "0" + date.getHours();
  let minutes = "0" + date.getMinutes();
  let seconds = "0" + date.getSeconds();
  let formattedDate = `${day.substr(-2)}.${month.substr(-2)}.${year}`;
  let formattedTime = `${hours.substr(-2)}:${minutes.substr(-2)}`;
  return `${formattedDate} ${formattedTime}`;
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
  if (validators[prop]) {
    return validators[prop](value);
  }
  return null;
}

export function toColor(num) {
  num >>>= 0;
  var b = num & 0xFF,
      g = (num & 0xFF00) >>> 8,
      r = (num & 0xFF0000) >>> 16;
  return "rgb(" + [r, g, b].join(",") + ")";
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

