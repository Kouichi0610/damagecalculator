import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types'
import { TargetSelectState, CandidatesFilter, Species } from './types'

export const actions: ActionTree<TargetSelectState, RootState> = {
  getCandidates: ({ commit }, filter: CandidatesFilter) => {
    axios.get('filtered_list', {
      params: {
        Types: filter.type,
        Total: filter.total,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      let sp = JSON.parse(json);
      let species: Species[] = new Array();
      for (var i = 0; i < sp.length; i++) {
        let s = new Species(sp[i]);
        species.push(s);
      }
      commit('setCandidates', species);
      return true;
    })
    .catch((e) => {
        console.log('failed:' + e);
        return false;
    });
  },
}

