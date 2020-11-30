import {writable} from 'svelte/store';
import {getUser} from './auth.js';

const projects = writable({});
const user = getUser();

export function getProjects(onUnsetCurrent) {
  let { subscribe, set, update } = projects;
  let prevUser = {};
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

  function setCurrent(id) {
    update((value) => {
      value.current = id;
    });
  }
  function unsetCurrent() {
    if (onUnsetCurrent) {
      onUnsetCurrent();
    }
    update((value) => {
      value.current = undefined;
    });
  }

  function release() {
    set({});
  }

  user.subscribe((value) => {
    if (!value.authorized) {
      release();
    }
    if (user != prevUser) {
      prevUser = user;
      refresh();
    }
  });

  return {
    subscribe,
    setCurrent,
    unsetCurrent,
    refresh
  };
}
