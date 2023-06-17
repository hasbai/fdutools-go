import { ref, computed, reactive } from "vue";
import { defineStore } from "pinia";
import {utils} from "../../wailsjs/go/models";

export const useStore = defineStore("counter", () => {
  const user = reactive(utils.User.createFrom({}));
  const setUser = (newUser: utils.User) => {
    user.name = newUser.name;
    user.uid = newUser.uid;
    user.pwd = newUser.pwd;
    user.profileID = newUser.profileID;
  }
  return { user, setUser };
});
