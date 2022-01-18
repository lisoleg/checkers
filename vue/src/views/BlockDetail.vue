<template>
  <div>
    <div class="container">
      <SpBlockDisplayFull :block="block"  tsFormat="MMM D YYYY, HH:mm:ss" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Block Detail',
  props:['blockHeight'],
  computed: {
    block() {
      return this.$store.getters['common/blocks/getBlockByHeight'](this.blockHeight)
    },
  },
}
</script>
