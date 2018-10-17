<template>
  <div
    class="relative z-10 user-button">
    <div>
      <div
        class="relative dropdown-trigger"
        tabindex="0"
        @click="toggle"
        @blur="toggle(false)"
        @keyup.enter="toggle"
      >
        <img
          v-if="user.photoURL"
          :src="user.photoURL"
          class="user-button__icon"
          aria-haspopup="true"
          aria-controls="dropdown-menu  ">
        <div
          v-else
          class="user-button__icon">
          <UserIcon />
        </div>
      </div>
        <div
            :class="{'hidden': !isActive}"
            class="absolute z-10 list-reset shadow-md mt-10 pin-t pin-r bg-white">
          <div class="p-4">
            <div v-if="!user.isAnonymous">
              <p v-if="user.displayName"><strong>{{ user.displayName }}</strong></p>
              <p>{{ user.email }}</p>
            </div>
            <div v-else>
              <p><em>Anonymous User</em></p>
            </div>
        </div>
          <div class="py-2 border-t border-grey">
            <a
              class="p-2 block w-full h-full cursor-pointer hover:bg-grey-lighter"
              tabindex="0"
              @blur="toggle(false)"
              @focus="toggle(true)"
              @click="signOut"
              @keyup.enter="signOut">
              Sign out
            </a>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import {mapActions, mapGetters} from 'vuex';
import UserIcon from 'vue-feather-icons/icons/UserIcon';
import isUndefined from 'lodash/isUndefined';

export default {
    components: {
        UserIcon
    },
    data() {
        return {
            isActive: false
        };
    },
    computed: {
        ...mapGetters({
            user: 'auth/user'
        })
    },
    methods: {
        ...mapActions({
            signOut: 'auth/signOut'
        }),
        toggle(override) {
            this.isActive = isUndefined(override) ? !this.isActive : override;
        }
    }
};
</script>
<style lang="scss" scoped>
.dropdown-trigger {
    height: 28px;
    width: 28px;
}

.user-button__icon {
    border-radius: 50%;
    cursor: pointer;
    height: 28px;
    width: 28px;
    color: #363636;
    background-color: white;

    > svg {
        height: 28px;
        width: 28px;
    }
}

.dropdown,
.dropdown-menu {
    // outline: none;
}
</style>
