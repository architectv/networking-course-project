import {writable} from 'svelte/store';
import {user} from './auth';
import {boards} from './boards';
import {lists} from './lists';

let task_labels_stores = {};

export function getTaskLabels(tasks, tid) {
  if (tasks == undefined || tid == undefined) {
    return undefined;
  }
  if (!task_labels_stores[taskId]) {
    task_labels_stores[taskId] = writable({});
  }
  let { subscribe, set, update } = task_labels_stores[taskId];
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
    let obj = await fetch(`api/v1/projects/${projectId}/boards/${boardId}` +
                          `/lists/${listId}/tasks/${taskId}/labels`, {
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
        value.error = "Load tasks error";
        return value;
      });
      console.log("error: ", x);
    });
  }

  async function addLabel(lid, onError) {
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
                              `/lists/${listId}/tasks/${taskId}/labels/${lid}`, {
      method: "POST",
      headers: {
        'Authorization': 'Bearer ' + token
      },
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

  async function deleteLabel(lid) {
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
                          `/lists/${listId}/tasks/${taskId}/labels/${lid}`, {
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

  let unsubscribe = tasks.subscribe((value) => {
    refresh();
  });

  function release() {
    set({});
    unsubscribe();
  }
  
  return {
    subscribe,
    listId,
    taskId,
    addLabel,
    deleteLabel,
    refresh,
    release,
  };
}
