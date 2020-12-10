import {writable} from 'svelte/store';
import {validate, validate_prop} from '../utils.js';

const user_store = writable({});

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

function getUser() {
  const { subscribe, set, update } = user_store;
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
  
    async function setUserdata() {
      let userdata = await fetch("api/v1/users", {
        method: "GET",
        headers: {
          Authorization: 'Bearer ' + localStorage.token
        }
      }).then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      }).then((x) => {
        return {authorized: true, data: x.data.user};
      }).catch((x) => {
        console.log("Login:", x);
        return {authorized: false, error: "Login error"};
      });
      console.log("Set data", userdata);
      set(userdata);
    }

    async function login(data) {
      let success = await fetch("api/v1/users/signin", {
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
      if (!success.authorized) {
        set(success)
      }
      await setUserdata();
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
          'Authorization': 'Bearer ' + token
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
      setUserdata();
    }

    return {
      subscribe,
      login,
      register,
      logout,
      unauthorized,
      validate: (data) => validate(validators, data),
      validate_prop: (prop, val) => validate_prop(validators, prop, val)
    };
  }

export const user = getUser();
