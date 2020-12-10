import {writable} from 'svelte/store';
import {user} from './auth';
import {boards} from './boards';
import {lists} from './lists';
import {validate, validate_prop} from '../utils.js';

let tasks_stores = [];

let validators = {
  title: (value) => {
    if (value.length <= 50) {
      return null;
    }
    return "Length of title > 50";
  }
};

export function getTasks(listId) {
  if (!tasks_stores[listId]) {
    tasks_stores[listId] = writable({});
  }
  let { subscribe, set, update } = tasks_stores[listId];
  async function refresh() {
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
      return;
    }
    let obj = await fetch("api/v1/projects/" + projectId + 
                          "/boards/" + boardId + "/lists/" +
                          listId + "/tasks", {
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
        value.list = x.data.tasks;
        value.error = undefined;
        return value;
      });
    }).catch((x) => {
      update((value) => {
        value.error = "Load tasks error";
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
    let boardId = lists.getBoardId();
    if (boardId == undefined) {
      return;
    }
    let token = localStorage.token;
    if (!token) {
      user.unauthorized();
      return;
    }
    let success = await fetch("api/v1/projects/" + projectId + 
                              "/boards/" + boardId + "/lists/" + 
                              listId + "/tasks", {
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

  let unsubscribe = boards.subscribe((value) => {
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

  async function deleteTask(id) {
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
    let success = await fetch(`api/v1/projects/${projectId}/boards/${boardId}` + 
                              `/lists/${listId}/tasks/${id}`, {
      method: "DELETE",
      headers: {
        'Authorization': 'Bearer ' + token
      },
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

  
  async function updateTask(id, prevListId, data) {
    console.log("UpdateTask", id, prevListId, data);
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
    let success = await fetch(`api/v1/projects/${projectId}/boards/${boardId}` + 
                              `/lists/${prevListId}/tasks/${id}`, {
      method: "PUT",
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
    unsubscribe();
    set({});
  }

  return {
    subscribe,
    listId,
    updateTask,
    deleteTask,
    refresh,
    create,
    release,
    validate: (data) => validate(validators, data),
    validate_prop: (prop, val) => validate_prop(validators, prop, val)
  };
}
