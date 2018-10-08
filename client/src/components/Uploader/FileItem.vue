<template>
  <div
    :class="[progress === 100 ? 'bg-green-lighter' : 'bg-grey-lighter']"
    class="w-full my-2 p-2 rounded relative">
    <span class="text-black relative z-10 text-sm">
      <FileIcon
        v-if="!progress"
        class="text-grey-darker h-4 w-4 text-center align-text-top" />
      <span
        v-else-if="progress === 100"
        class="inline-block bg-green rounded-full h-4 w-4">
        <CheckIcon class="check text-white h-3 w-4 align-middle font-semibold" />
      </span>
      <ClockIcon
        v-else
        class="text-grey-darker h-4 w-4 align-middle hover:text-black align-text-bottom" />
      {{ file.name }}
      <XIcon
        @click="onRemove"
        v-if="!progress"
        class="text-grey-darker h-4 w-4 align-middle float-right hover:text-black align-text-bottom" />
    </span>
    <div
      v-if="progress !== 100"
      :style="{'width': progress + '%'}"
      class="transition-width absolute h-full bg-grey pin-t pin-l" />
  </div>
</template>
<script>
import FileIcon from 'vue-feather-icons/icons/FileIcon';
import XIcon from 'vue-feather-icons/icons/XIcon';
import CheckIcon from 'vue-feather-icons/icons/CheckIcon';
import ClockIcon from 'vue-feather-icons/icons/ClockIcon';
export default {
    components: {
        FileIcon,
        XIcon,
        CheckIcon,
        ClockIcon
    },
    props: {
        file: {
            type: File,
            required: true
        },
        index: {
            type: Number
        },
        progress: {
            type: Number,
            default: 0
        }
    },
    methods: {
        onRemove(e) {
            this.$emit('remove', this.index);
        }
    }
};
</script>
<style scoped>
.check {
    stroke-width: 3;
}

.transition-width {
    transition: width 200ms;
}
</style>
