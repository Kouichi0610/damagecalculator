// index.ts
// モジュール定義
import {Module} from 'vuex';
import { RootState } from '@/store/types'
import { TodosState } from '@/store/types'
import getters from './getters';
import actions from './actions'
import mutations from './mutations'

export const state: TodosState = {
    title: 'todoList',
    todos: [],
};

export const todos: Module<TodosState, RootState> = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};
