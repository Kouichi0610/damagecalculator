<template>
    <div class="todo">
      TODO数:{{ size }}<br>
      <!-- first {{ first }} <br> -->
      <p>Current{{ current }}</p>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import {State, Action, Getter} from 'vuex-class';
import Component from 'vue-class-component';
// TODO:使っているはずなのにエラーが出る
//import {Todo, TodosState} from '../store/todo/types'
const namespace: string = 'todos';

@Component
export default class TodoComponent extends Vue {
  @State('todos') todos: TodosState;

  @Getter('size', { namespace })
  private size!: number;

  @Getter('first', { namespace })
  private first!: () => string;

  @Action('add', { namespace })
  private add!: (todo: Todo) => Promise<boolean>;
  @Action('remove', { namespace })
  private remove!: (id: string) => Promise<boolean>;

  private inputText: string = '';


  mounted() {
    this.add({
      id: '15',
      done: false,
      text: 'abcd',
    });
    this.add({
      id: '20',
      done: false,
      text: 'xxxx',
    });
    this.remove('15');
    // console.log('' + this.todos[0].text);
    //console.log('x ' + this.todoState.todos[0].text);
    //this.todoState.title = 'abcd';
    console.log('title:' + this.todos.title);
  }

  get current(): string {
    // TODO:直接todos[0].textと書くとエラーが出る
    const first: Todo = this.todos && this.todos.todos && this.todos.todos[0];
    return (first && first.text) || '';
  }

  private addTodo(): void {
    // validate
    if (!this.inputText) {
      window.alert('なにか入力してね！');
      return;
    }
    this.add({
      id: new Date().getMilliseconds().toString(),
      // チェックボックスON/OFF
      done: false,
      // やること
      text: this.inputText,
    });
  }
}
</script>
