import {writable} from 'svelte/store';
import {user} from './auth.js';
import {boards} from './boards.js';
import {lists} from './lists';
import {validate, validate_prop} from '../utils.js';

const labels_store = writable({});

let validators = {
  name: (value) => {
    if (value.length <= 50) {
      return null;
    }
    return "Length of title > 50";
  }
};

function getLabels() {
  let { subscribe, set, update } = labels_store;
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
                          "/boards/" + boardId + "/labels", {
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
        value.list = x.data.labels;
        value.error = undefined;
        return value;
      });
    }).catch((x) => {
      update((value) => {
        value.error = "Load labels error";
        return value;
      });
      console.log("error: ", x);
    });
  }
  
  async function deleteLabel(id) {
    let projectId = boards.getProjectId();
    if (projectId == undefined) {
      return;
    }
    let boardId = lists.getBoardId();
    if (boardId == undefined) {
      return;
    }
    let token = localStorage.token;
    if (!token) {
      user.unauthorized();
      return;
    }
    let success = await fetch(`api/v1/projects/${projectId}/boards/${boardId}/labels/${id}`, {
      method: "DELETE",
      headers: {
        'Authorization': 'Bearer ' + token
      }
    }).then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      return true;
    }).catch((x) => {
      return false;
    });
    if (success) {
      refresh();
      lists.fullReload();
    }
    
  }

  async function create(data, onError) {
    if (data.color == "") {
      data.color = 0;
    } else {
      data.color = parseInt(data.color.substr(1), 16);
    }
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
                              "/boards/" + boardId + "/labels", {
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

  async function updateLabel(id, data) {
    let projectId = boards.getProjectId();
    if (projectId == undefined) {
      return;
    }
    let boardId = lists.getBoardId();
    if (boardId == undefined) {
      return;
    }
    let token = localStorage.token;
    if (!token) {
      user.unauthorized();
      return;
    }
    let success = await fetch(`api/v1/projects/${projectId}/boards/${boardId}/labels/${id}`, {
      method: "PATCH",
      headers: {
        'Content-Type': 'application/json;charset=utf-8',
        'Authorization': 'Bearer ' + token
      },
      body: JSON.stringify(data)
    }).then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    }).then((x) => {
      return true;
    }).catch((x) => {
      return false;
    });
    if (success) {
      refresh();
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
    release,
    updateLabel,
    deleteLabel,
    validate: (data) => validate(validators, data),
    validate_prop: (prop, val) => validate_prop(validators, prop, val)
  };
}

export const labels = getLabels();
