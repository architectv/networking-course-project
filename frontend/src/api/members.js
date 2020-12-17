import { writable } from "svelte/store";
import { user } from "./auth";

export function getMembers(prefix) {
  let { subscribe, set, update } = writable([]);
  async function refresh() {
    if (!localStorage.token) {
      return [];
    }
    await fetch(`${prefix}/members`, {
      headers: {
        Authorization: `Bearer ${localStorage.token}`
      }
    }).then((resp) => {
      if (resp.status == 401) {
        user.unauthorized();
      }
      if (resp.ok) {
        return resp.json();
      }
      throw Error("Network error");
    }).then((data) => {
      let m = data.data.members;
      let owner = [];
      let admins = [];
      let writers = [];
      let readers = [];
      m.forEach((value) => {
        if (value.isOwner) owner.push(value);
        else if (value.permissions.admin) admins.push(value);
        else if (value.permissions.write) writers.push(value);
        else readers.push(value);
      });
      let res = [];
      res.push({name: 'Бог', items: owner});
      if (admins.length) res.push({name: 'Админы', items: admins});
      if (writers.length) res.push({name: 'Писатели', items: writers});
      if (readers.length) res.push({name: 'Читатели', items: readers});
      set(res);
    }).catch((e) => {
      console.log(e);
      set([]);
    });
  }
  
  async function addMember(member_name, role) {
    let data = {read: true};
    if (role == 'writer' || role == 'admin') {
      data.write = true;
    }
    if (role == 'admin') {
      data.admin = true;
    }
    if (!member_name) {
      return;
    }
    if (!localStorage.token) {
      return [];
    }
    await fetch(`${prefix}/permissions?member_nickname=${member_name}`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        'Content-Type': 'application/json;charset=utf-8',
        Authorization: `Bearer ${localStorage.token}`
      }
    }).then((resp) => {
      if (response.status == 401) {
        user.unauthorized();
      }
      if (resp.ok) {
        return resp.json();
      }
      throw Error("Network error");
    }).then((data) => {
      refresh();
    }).catch((e) => {
    });
  }

  refresh();
  return {
    refresh,
    update,
    addMember,
    set,
    subscribe
  }
}