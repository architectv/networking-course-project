import {writable} from 'svelte/store';
import {user} from './auth.js';
import {validate, validate_prop} from '../utils.js';

const projects_store = writable({});

let validators = {
  title: (value) => {
    if (value.length <= 50) {
      return null;
    }
    return "Length of title > 50";
  }
};

function getProjects() {
  let { subscribe, set, update } = projects_store;
  async function refresh() {
    let token = localStorage.token;
    if (!token) {
      return;
    }
    let obj = await fetch("api/v1/projects", {
      headers: {
        Authorization: "Bearer " + token
      },
    }).then((response) => {
      if (response.status == 401) {
        user.unauthorized();
      }
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      set({list: x.data.projects});
    }).catch((x) => {
      update((value) => {
        value.error = "Load projects error";
        return value;
      });
      console.log("error: ", x);
    });
  }
  
  async function deleteCurrent(onError) {
    let token = localStorage.getItem("token");
    if (!token) {
      user.unauthorized();
      return;
    }
    let current;
    update((value) => {
      current = value.current;
      return value;
    })
    if (!current) {
      return;
    }
    let success = await fetch("api/v1/projects/" + current, {
      method: "DELETE",
      headers: {
        'Authorization': 'Bearer ' + token
      }
    }).then((response) => {
      if (response.status == 401) {
        user.unauthorized();
        throw new Error('Unauthorized');
      }
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      return true;
    }).catch((x) => {
      if (onError) onError(x);
      return false;
    });
    if (success) {
      unsetCurrent();
      await refresh();
    }
  }

  async function updateCurrent(data, onError) {
    let p = {};
    p.read = data.read;
    data.read = undefined;
    p.write = data.write;
    data.write = undefined;
    data.defaultPermissions = p;
    let token = localStorage.getItem("token");
    if (!token) {
      user.unauthorized();
      return;
    }
    let current;
    update((value) => {
      current = value.current;
      return value;
    })
    if (!current) {
      return;
    }
    let success = await fetch(`api/v1/projects/${current}`, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: {
        'Content-Type': 'application/json;charset=utf-8',
        'Authorization': 'Bearer ' + token
      }
    }).then((response) => {
      if (response.status == 401) {
        user.unauthorized();
        throw new Error('Unauthorized');
      }
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      return true;
    }).catch((x) => {
      if (onError) onError(x);
      return false;
    });
    if (success) {
      await refresh();
    }
    return success;
  }

  async function create(data, onError) {
    let token = localStorage.getItem("token");
    if (!token) {
      user.unauthorized();
      return;
    }
    let success = await fetch("api/v1/projects", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json;charset=utf-8',
        'Authorization': 'Bearer ' + token
      },
      body: JSON.stringify(data)
    }).then((response) => {
      if (response.status == 409) {
        throw new Error('Already exists');
      }
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      return true;
    }).catch((x) => {
      if (onError) onError(x);
      return false;
    });
    if (success) {
      await refresh();
    }
  }

  function setCurrent(id) {
    update((value) => {
      value.current = id;
      return value;
    });
  }
  function unsetCurrent() {
    update((value) => {
      value.current = undefined;
      return value;
    });
  }

  function release() {
    set({});
  }

  user.subscribe((value) => {
    if (!value.authorized) {
      release();
    } else {
      refresh();
    }
  });

  return {
    subscribe,
    setCurrent,
    unsetCurrent,
    refresh,
    release,
    deleteCurrent,
    updateCurrent,
    create,
    validate: (data) => validate(validators, data),
    validate_prop: (prop, val) => validate_prop(validators, prop, val)
  };
}

export const projects = getProjects();
