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

export function getTasks(pid, bid, listId) {
  if (!tasks_stores[listId]) {
    tasks_stores[listId] = writable({});
  }
  let { subscribe, set, update } = tasks_stores[listId];
  async function refresh() {
    let token = localStorage.token;
    if (bid == undefined || pid == undefined || listId == undefined || !token) {
      return;
    }
    let obj = await fetch(`api/v1/projects/${pid}/boards/${bid}/lists/${listId}/tasks`, {
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
    refreshLabels();
  }
  
  async function updateLabel(method, tid, lid) {
    let token = localStorage.token;
    if (bid == undefined || pid == undefined || tid == undefined || !method ||
        listId == undefined || lid == undefined || !token) {
      return;
    }
    return await fetch(`api/v1/projects/${pid}/boards/${bid}/` +
                          `lists/${listId}/tasks/${tid}/labels/${lid}`, {
      method,
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
      loadLabels(tid);
    }).catch((x) => {
      console.log("error: ", x);
    });
  }
  
  async function addLabel(tid, lid) {
    await updateLabel("POST", tid, lid);
  }
  
  async function removeLabel(tid, lid) {
    await updateLabel("DELETE", tid, lid);
  }
  
  async function loadLabels(tid) {
    let token = localStorage.token;
    if (bid == undefined || pid == undefined || tid == undefined ||
        listId == undefined || !token) {
      return;
    }
    return await fetch(`api/v1/projects/${pid}/boards/${bid}/` +
                          `lists/${listId}/tasks/${tid}/labels`, {
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
      let labels = x.data.labels;
      update((value) => {
        if (value.list) {
          value.list.forEach((task) => {
            if (task._id == tid) {
              if (!labels) {
                task.labels = [];
              } else {
                task.labels = labels;
              }
            }
          });
        }
        return value;
      });
    }).catch((x) => {
      console.log("error: ", x);
    });
  }
  
  async function refreshLabels() {
    let tlist;
    update((value) => {
      tlist = value.list;
      return value;
    });
    if (!tlist) {
      return;
    }
    for (const value of tlist) {
      await loadLabels(value._id);
    }
  }

  async function create(data, onError) {
    let token = localStorage.token;
    if (bid == undefined || pid == undefined || listId == undefined || !token) {
      return;
    }
    let success = await fetch("api/v1/projects/" + pid + 
                              "/boards/" + bid + "/lists/" + 
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

  async function deleteTask(id) {
    let token = localStorage.token;
    if (bid == undefined || pid == undefined || listId == undefined || id == undefined || !token) {
      return;
    }
    let success = await fetch(`api/v1/projects/${pid}/boards/${bid}` + 
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

  let unsubscribe = lists.subscribe((value) => {
    console.log("List", value);
    if (!value.list) {
      set({});
    } else {
      refresh();
    }
  });
  

  
  async function updateTask(id, prevListId, data) {
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
      set({});
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
    addLabel,
    removeLabel,
    validate: (data) => validate(validators, data),
    validate_prop: (prop, val) => validate_prop(validators, prop, val)
  };
}
