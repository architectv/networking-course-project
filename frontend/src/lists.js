import {writable} from 'svelte/store';
import {user} from './auth.js';
import {projects} from './projects.js';
import {boards} from './boards.js';
import {validate, validate_prop} from './utils.js';

const lists_store = writable({});

let validators = {
  title: (value) => {
    if (value.length <= 50) {
      return null;
    }
    return "Length of title > 50";
  }
};

function getLists() {
  let { subscribe, set, update } = lists_store;
  function getBoardId() {
    let boardId;
    update((value) => {
      boardId = value.boardId;
      return value;
    });
    return boardId;
  }
  async function refresh() {
    let projectId = boards.getProjectId();
    if (projectId == undefined) {
      return;
    }
    let boardId = getBoardId();
    if (boardId == undefined) {
      return;
    }
    let token = localStorage.token;
    if (!token) {
      return;
    }
    let obj = await fetch("api/v1/projects/" + projectId + 
                          "/boards/" + boardId + "/lists", {
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
      update((value) => {
        value.list = x.data.lists;
        value.error = undefined;
        return value;
      });
    }).catch((x) => {
      update((value) => {
        value.error = "Load lists error";
        return value;
      });
      console.log("error: ", x);
    });
  }

  async function create(data, onError) {
    let projectId = boards.getProjectId();
    if (projectId == undefined) {
      return;
    }
    let boardId = getBoardId();
    if (boardId == undefined) {
      return;
    }
    let token = localStorage.token;
    if (!token) {
      user.unauthorized();
      return;
    }
    let success = await fetch("api/v1/projects/" + projectId + 
                              "/boards/" + boardId + "/lists", {
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

  function release() {
    set({});
  }

  boards.subscribe((value) => {
    let newBoardId = value.current;
    let boardId;
    update((value) => {
      boardId = value.boardId;
      if (!newBoardId) {
        return {};
      }
      value.boardId = newBoardId;
      return value;
    });
    if (newBoardId != boardId) {
      refresh();
    }
  });

  return {
    getBoardId,
    subscribe,
    refresh,
    create,
    validate: (data) => validate(validators, data),
    validate_prop: (prop, val) => validate_prop(validators, prop, val)
  };
}

export const lists = getLists();
