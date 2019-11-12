<script>
  export default {
    render (h, cxt) {
      return h(require(`./forms/${this.formName}.vue`), {
        props: {
          prop: this.prop
        }
      })
    },
    props: {
      formName: String,
      prop: Object
    }
  }
</script>