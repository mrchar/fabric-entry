<template>
  <div class="container">
    <a-form :form="form">
      <a-form-item label="组织名称">
        <a-input v-model="form.orgName" />
      </a-form-item>
    </a-form>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component'

@Options({
  data () {
    return {
      form: {
        orgName: ''
      }
    }
  }
})
export default class QuickStart extends Vue {}
</script>

<style scoped>
  .container {
    width: 600px;
    margin: 20px auto;
  }
</style>
