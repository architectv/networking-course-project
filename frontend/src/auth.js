import {getCookie, setCookie} from './utils';
import {writable} from 'svelte/store';

const user = writable({});

export function getUser() {
	const { subscribe, set, update } = user;

	return {
		subscribe,
    login: async (data) => {
      console.log("login");
      let userdata = await fetch("api/auth/users", {
      }).then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      }).then((x) => {
        x = {resp: x};
        x.authorized = true;
        x.username = data.username;
        return x;
      }).catch((x) => {
        return {};
      });
      set(userdata);
    },
    register: (data) => {
      data.authorized = true;
      set(data);
    },
    logout: () => {
      set({});
    },
	};
}
