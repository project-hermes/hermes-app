<template>
  <div
    class="relative z-10">
    <div>
      <div
        class="h-10 w-10 rounded-full border-4 border-blue hover:border-blue-dark"
        tabindex="0"
        @click="toggle()"
        @keyup.enter="toggle()"
        ref="button"
      >
        <img
          v-if="user.photoURL"
          :src="user.photoURL"
          class="h-8 w-8 cursor-pointer text-white bg-white rounded-full border-2 border-blue-light"
          :class="{'border-white': isActive || $route.name === 'accountSettings'}"
          aria-haspopup="true"
          aria-controls="dropdown-menu">
        <div
          v-else
          class="user-button__icon">
          <UserIcon />
        </div>
      </div>
        <div
            v-closable="{
                exclude: ['button'],
                handler: close
            }"
            v-if="isActive"
            class="absolute z-10 list-reset shadow-md mt-14 pin-t pin-r bg-white text-lg">
          <div class="p-6">
            <div>
              <p v-if="user.displayName" class="font-semibold">
                  {{ user.displayName }}
              </p>
              <p>
                  {{ user.email }}
              </p>
            </div>
        </div>
          <div class="py-2 border-t border-grey">
            <a
                class="py-2 px-6 block w-full h-full cursor-pointer hover:bg-grey-lighter text-black no-underline"
                tabindex="0"
                @click="goToAccountSettings"
                @keyup.enter="goToAccountSettings">
                Account Settings
            </a>
        </div>
        <div class="py-2 border-t border-grey">
          <a
            class="py-2 px-6 block w-full h-full cursor-pointer hover:bg-grey-lighter"
            tabindex="0"
            @click="signOut"
            @keyup.enter="signOut">
            Sign Out
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
        toggle (override) {
            this.isActive = isUndefined(override) ? !this.isActive : override;
        },
        close () {
            this.isActive = false;
        },
        goToAccountSettings () {
            this.isActive = false;
            this.$router.push({name: 'accountSettings'});
        }
    }
};
</script>
