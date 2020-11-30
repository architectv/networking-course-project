import {writable} from 'svelte/store';

const user = writable({});

let validators = {
  nickname: (value) => {
    if (value) {
      if (value.length >= 3 && value.length <= 32) {
        return null;
      }
      return "Length must be in range (3, 32)";
    }
    return "Empty field";
  },
  password: (value) => {
    if (value) {
      if (value.length >= 6 && value.length <= 32) {
        return null;
      }
      return "Length must be in range (6, 32)";
    }
    return "Empty field";
  },
  email: (value) => {
    if (/^[a-zA-Z0-9.!#$%&*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(value)) {
      return null;
    }
    return "Invalid email";
  }
};

export function getUser() {
  const { subscribe, set, update } = user;
  function validate_prop(prop, value) {
    console.log("Validate", prop, value)
    if (validators[prop]) {
      return validators[prop](value);
    }
    return null;
  }

  function validate(data) {
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

    async function register(data) {
      let success = await fetch("api/v1/users/signup", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
      }).then((response) => {
        if (response.status == 409) {
          set({authorized: false, error: "Already exists"});
        }
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      }).then((x) => {
        return true;
      }).catch((x) => {
        console.log("error: ", x);
        return false;
      });
      if (success) {
        await login({nickname: data.nickname, password: data.password});
      }
    }

    async function login(data) {
      let userdata = await fetch("api/v1/users/signin", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
      }).then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      }).then((x) => {
        console.log("Login: ", x);
        localStorage.setItem("token", x.data.token);
        return {authorized: true, token: x.data.token, nickname: data.nickname};
      }).catch((x) => {
        console.log("Login:", x);
        return {authorized: false, error: "Login error"};
      });
      console.log("Set data", userdata);
      set(userdata);
    }

    async function logout() {
      let token = localStorage.getItem("token");
      if (!token) {
        console.log("Logout: empty token");
        return;
      }
      let success = await fetch("api/v1/users/signout", {
        method: "GET",
        headers: {
          'Authorization': 'Bearer ' + localStorage.getItem("token")
        }
      }).then((response) => {
        if (response.status == 401) {
          unauthorized();
          return;
        }
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      }).then((x) => {
        localStorage.removeItem("token");
        return true;
      }).catch((x) => {
        console.log("error: ", x);
        return false;
      });
      if (success) {
        set({authorized: false});
      }
    }
    function unauthorized() {
      localStorage.removeItem("token");
      set({authorized: false, error: "Unauthorized"});
    }
    
    if (localStorage.token) {
      set({authorized: true, token: localStorage.token});
    }

    return {
      subscribe,
      validate,
      login,
      register,
      logout,
      validate_prop,
      unauthorized
    };
  }
