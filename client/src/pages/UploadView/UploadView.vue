<template>
  <main class="upload-view">
    <div class="columns is-centered">
      <div class="upload-view--container">
        <div class="field">
          <div class="file is-medium is-boxed has-name is-primary">
            <label class="file-label">
              <input
                class="file-input"
                type="file"
                ref="upload"
                @change="onFileSelect">
              <span class="file-cta">
                <span class="file-icon">
                  <UploadCloudIcon />
                </span>
                <span class="file-label">
                  Select a dive
                </span>
              </span>
              <span
                :class="{'upload-view--file-name--empty': !fileName}"
                class="upload-view--file-name file-name">
                {{ fileName }}
              </span>
            </label>
          </div>
        </div>
        <p v-if="confirming" class="is-centered buttons">
          <a @click="onConfirmation" class="button is-success">
            <span class="icon is-small">
                <CheckIcon />
            </span>
            <span>Upload</span>
          </a>
          <a @click="onCancel" class="button is-danger is-outlined">
            <span>Cancel</span>
          </a>
        </p>
        <progress
          v-if="status === 'uploading'"
          :value="progress"
          class="progress is-success"
          max="100"/>
        <div v-else-if="status === 'complete'">
            <h4 class="title is-4 has-text-centered">Thanks!</h4>
        </div>
        <div class="is-centered" v-else-if="status === 'error'">
            <div>Oops. Something went wrong...</div>
            <p class="is-centered buttons">
                <a @click="onConfirmation" class="button is-success is-centered">
                  <span>Try again?</span>
                </a>
            </p>
        </div>
      </div>
    </div>
  </main>
</template>
<script>
import UploadCloudIcon from 'vue-feather-icons/icons/UploadCloudIcon';
import CheckIcon from 'vue-feather-icons/icons/CheckIcon';
import head from 'lodash/head';
import {mapGetters, mapActions, mapMutations} from 'vuex';

export default {
    components: {
        UploadCloudIcon,
        CheckIcon
    },
    data() {
        return {
            file: null,
            confirming: false
        };
    },
    computed: {
        ...mapGetters({
            progress: 'upload/progress',
            status: 'upload/status'
        }),
        fileName() {
            return this.file ? this.file.name : '';
        }
    },
    methods: {
        ...mapActions({
            uploadFile: 'upload/uploadFile'
        }),
        ...mapMutations({
            resetStatus: 'upload/resetStatus'
        }),
        onFileSelect(event) {
            this.resetStatus();
            this.file = event.target.files.item(0);
            this.confirming = !!this.file;
        },
        onCancel() {
            this.confirming = false;
            this.file = null;
            this.$refs.upload.value = '';
        },
        onConfirmation() {
            this.confirming = false;
            this.uploadFile({
                file: this.file
            }).then(() => {
                this.$refs.upload.value = '';
            });
        }
    }
};
</script>
<style lang="scss">
.upload-view {
    height: 100%;

    &--container {
        margin-top: 5rem;
    }
}
</style>
