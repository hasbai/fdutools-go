import { ref, computed, reactive } from "vue";
import { defineStore } from "pinia";
import { utils, xk } from "../../wailsjs/go/models";

export const useStore = defineStore("counter", () => {
  const user = reactive(utils.User.createFrom({}));
  const setUser = (newUser: utils.User) => {
    user.name = newUser.name;
    user.uid = newUser.uid;
    user.pwd = newUser.pwd;
    user.profileID = newUser.profileID;
  };

  const currentCourse = ref(new xk.Course());
  const setCurrentCourse = (value: xk.Course) => {
    if (!value.name) {
      value.name = currentCourse.value.name;
      setTimeout(() => {
        value.name = "";
      }, 150);
    }
    currentCourse.value = value;
  };
  return { user, setUser, currentCourse, setCurrentCourse };
});
