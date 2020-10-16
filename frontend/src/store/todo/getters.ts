import { GetterTree } from 'vuex';
import { TodosState } from './types';
import { RootState } from '@/store/types';

const getters: GetterTree<TodosState, RootState> = {
  size: (state: TodosState) => {
    return state.todos.length;
  },
  first: (state: TodosState) => {
      return state.todos[0].text;
  }

};

export default getters;