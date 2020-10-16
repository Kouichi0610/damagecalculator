import { MutationTree } from 'vuex';
import { TodosState } from './types';
import { RootState } from '@/store/types';

const mutations: MutationTree<TodosState> = {
  add: (state, todo: Todo) => {
    state.todos.push(todo);
  },
  remove: (state, id: string) => {
    state.todos = state.todos.filter((e: Todo) => e.id !== id);
  },
};

export default mutations;