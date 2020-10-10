// 種族値
export const species = {
  namespaced: true,
  state: {
    hp: 95,
    at: 109,
    df: 105,
    sa: 75,
    sd: 85,
    sp: 56,
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
