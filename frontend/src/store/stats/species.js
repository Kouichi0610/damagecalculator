// 種族値
export const species = {
  namespaced: true,
  state: {
    hp: 10,
    at: 10,
    df: 10,
    sa: 10,
    sd: 10,
    sp: 10,
  },
  getters: {
    toString (state) {
      return '種族値:' + state.hp + ' ' + state.at + ' ' + state.df + ' ' + state.sa + ' ' + state.sd + ' ' + state.sp;
    }
  },
  
  mutations: {
    setSpecies (state, hp, at, df, sa, sd, sp) {
      state.hp = hp;
      state.at = at;
      state.df = df;
      state.sa = sa;
      state.sd = sd;
      state.sp = sp;
    }
  }
}
