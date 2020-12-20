<template>
  <q-page class="flex flex-center">
<!--    <img alt="Quasar logo" src="../assets/logo.png">-->

    <transition
            appear
            enter-active-class="animated fadeIn"
            leave-active-class="animated fadeOut"
    >
      <!-- Wrapping only one DOM element, defined by QBtn -->
      <q-btn
              color="secondary"
              icon="mail"
              label="Email"
      />
    </transition>

    <template v-for="c in dora_components">
      <component  v-if="c.component_name == 'q-input'" :is="c.component_name" :label="c.label" :key="c.component_name" ></component>
    </template>

      <draggable >
          <transition-group>
              <template v-for="c in dora_components">
                  <component  v-if="c.component_name == 'q-input'" :is="c.component_name" :label="c.label" :key="c.component_name" ></component>
              </template>
          </transition-group>
      </draggable>
<!--    <li v-for="c in dora_components" :key="c">-->
<!--      {{c}}-->
<!--    </li>-->
  </q-page>
</template>

<style>
</style>

<script>
    import draggable from 'vuedraggable'
  let qinput = {
    component_name:'q-input',
    label:"姓名"
  }

  let qinput2 = {
    component_name:'q-input',
    label:"年龄"
  }
export default {
  name: 'HelloWorld',
    components: {
        draggable
    },
  data(){
    return {
      dora_components:[qinput,qinput2]
    }

  }

}
</script>
