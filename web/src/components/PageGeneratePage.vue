<template>
    <div class="q-pa-md">
        <div class="row">
            <div class="col">
                <q-list bordered padding class="rounded-borders" style="max-width: 350px">
                    <q-item-label header>表单组件</q-item-label>
                    <draggable>
                        <q-item v-ripple clickable v-on:click="dropComponent" >
                            <q-item-section avatar top >
                                <q-avatar icon="i" color="primary" text-color="white"/>
                            </q-item-section>

                            <q-item-section>
                                <q-item-label lines="1">文本框</q-item-label>
                                <q-item-label caption>文本输入框</q-item-label>
                            </q-item-section>
                        </q-item>
                    </draggable>

                    <q-item clickable v-ripple>
                        <q-item-section avatar top>
                            <q-avatar icon="folder" color="orange" text-color="white"/>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label lines="1">Movies</q-item-label>
                            <q-item-label caption>March 1st, 2019</q-item-label>
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="info"/>
                        </q-item-section>
                    </q-item>

                    <q-item clickable v-ripple>
                        <q-item-section avatar top>
                            <q-avatar icon="folder" color="teal" text-color="white"/>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label lines="1">Photos</q-item-label>
                            <q-item-label caption>January 15th, 2019</q-item-label>
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="info"/>
                        </q-item-section>
                    </q-item>

                    <q-separator spaced/>
                    <q-item-label header>Files</q-item-label>

                    <q-item clickable v-ripple>
                        <q-item-section avatar top>
                            <q-avatar icon="assignment" color="grey" text-color="white"/>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label lines="1">Expenses spreadsheet</q-item-label>
                            <q-item-label caption>March 2nd, 2019</q-item-label>
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="info"/>
                        </q-item-section>
                    </q-item>

                    <q-item clickable v-ripple>
                        <q-item-section avatar top>
                            <q-avatar icon="place" color="grey" text-color="white"/>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label lines="1">Places to visit</q-item-label>
                            <q-item-label caption>February 22, 2019</q-item-label>
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="info" color="amber"/>
                        </q-item-section>
                    </q-item>

                    <q-item clickable v-ripple>
                        <q-item-section avatar top>
                            <q-avatar icon="library_music" color="grey" text-color="white"/>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label lines="1">My favorite song</q-item-label>
                            <q-item-label caption>Singing it all day</q-item-label>
                        </q-item-section>

                        <q-item-section side>
                            <q-icon name="info"/>
                        </q-item-section>
                    </q-item>
                </q-list>
            </div>
            <div class="col-8">
                <draggable v-model="dora_components">
                    <template v-for="c in dora_components">
                        <component  v-if="c.component_name == 'q-input'" :is="c.component_name" :label="c.label" :key="c.component_name" @click="editProterty(c)" ></component>
                    </template>
                </draggable>
            </div>
            <div class="col">
                {{dora_components}}
<!--                <q-input v-model="current_select_comp.component_name" label="组件名"></q-input>-->
                <q-select v-model="current_select_comp.component_name" :options="options"></q-select>
                <q-input v-model="current_select_comp.label" label="label" ></q-input>
            </div>
        </div>
    </div>

</template>

<style>
</style>

<script>
      import draggable from 'vuedraggable'
    // let qinput = {
    //   component_name:'q-input',
    //   label:"姓名"
    // }
    //
    // let qinput2 = {
    //   component_name:'q-input',
    //   label:"年龄"
    // }
    export default {
        name: 'PageGeneatePage',
        components: {
            draggable
        },
        data() {
            return {
                dora_components:[],
                options:['q-input','q-select','q-color'],
                current_select_comp:{}
            }

        },
        methods:{
            dropComponent:function(e){
                console.log(e)
                let c = {
                    component_name:'q-input',
                    label:"姓名"
                }
                this.dora_components.push(c)
            },
            editProterty:function (comp_ref) {
                this.current_select_comp = comp_ref
                console.log(comp_ref)

            }
        }

    }
</script>
