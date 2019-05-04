<template>
  <div
    class="relative z-10">
    <div>
      <div
        class="h-10 w-10 rounded-full border-4 border-blue hover:border-blue-dark"
        :class="{'border-blue-dark': isActive}"
        tabindex="0"
        @click="toggle()"
        @blur="toggle(false)"
        @keyup.enter="toggle()"
      >
        <img
          v-if="user.photoURL"
          :src="user.photoURL"
          class="user-icon h-8 w-8 cursor-pointer text-white bg-white rounded-full border-2 border-blue-light"
          aria-haspopup="true"
          aria-controls="dropdown-menu">
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
              @click="signOut"
              @focus="toggle(true)"
              @blur="toggle(false)"
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
            setTimeout(() => {
                this.isActive = isUndefined(override) ? !this.isActive : override;
            });
        }
    }
};
</script>
