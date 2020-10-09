// 性格
/*
  TODO:設計
  テストをできれば

*/
const bashful = {
  name: 'てれや',
  at: 1.0,
  df: 1.0,
  sa: 1.0,
  sd: 1.0,
  sp: 1.0,
}
const lonely = {
  name: 'さみしがり',
  at: 1.1,
  df: 0.9,
  sa: 1.0,
  sd: 1.0,
  sp: 1.0,
}
const adamant = {
  name: 'いじっぱり',
  at: 1.1,
  df: 1.0,
  sa: 0.9,
  sd: 1.0,
  sp: 1.0,
}
const naughty = {
  name: 'やんちゃ',
  at: 1.1,
  df: 1.0,
  sa: 1.0,
  sd: 0.9,
  sp: 1.0,
}
const brave = {
  name: 'ゆうかん',
  at: 1.1,
  df: 1.0,
  sa: 1.0,
  sd: 1.0,
  sp: 0.9,
}
const bold = {
  name: 'ずぶとい',
  at: 0.9,
  df: 1.1,
  sa: 1.0,
  sd: 1.0,
  sp: 1.0,
}
const impish = {
  name: 'わんぱく',
  at: 1.0,
  df: 1.1,
  sa: 0.9,
  sd: 1.0,
  sp: 1.0,
}
const lax = {
  name: 'のうてんき',
  at: 1.0,
  df: 1.1,
  sa: 1.0,
  sd: 0.9,
  sp: 1.0,
}
const relaxed = {
  name: 'のんき',
  at: 1.0,
  df: 1.1,
  sa: 1.0,
  sd: 1.0,
  sp: 0.9,
}
const modest = {
  name: 'ひかえめ',
  at: 0.9,
  df: 1.0,
  sa: 1.1,
  sd: 1.0,
  sp: 1.0,
}
const mild = {
  name: 'おっとり',
  at: 1.0,
  df: 0.9,
  sa: 1.1,
  sd: 1.0,
  sp: 1.0,
}
const rash = {
  name: 'うっかりや',
  at: 1.0,
  df: 1.0,
  sa: 1.1,
  sd: 0.9,
  sp: 1.0,
}
const quiet = {
  name: 'れいせい',
  at: 1.0,
  df: 1.0,
  sa: 1.1,
  sd: 1.0,
  sp: 0.9,
}
const calm = {
  name: 'おだやか',
  at: 0.9,
  df: 1.0,
  sa: 1.0,
  sd: 1.1,
  sp: 1.0,
}
const gentle = {
  name: 'おとなしい',
  at: 1.0,
  df: 0.9,
  sa: 1.0,
  sd: 1.1,
  sp: 1.0,
}
const careful = {
  name: 'しんちょう',
  at: 1.0,
  df: 1.0,
  sa: 0.9,
  sd: 1.1,
  sp: 1.0,
}
const sassy = {
  name: 'なまいき',
  at: 1.0,
  df: 1.0,
  sa: 1.0,
  sd: 1.1,
  sp: 0.9,
}
const timid = {
  name: 'おくびょう',
  at: 0.9,
  df: 1.0,
  sa: 1.0,
  sd: 1.0,
  sp: 1.1,
}
const hasty = {
  name: 'せっかち',
  at: 1.0,
  df: 0.9,
  sa: 1.0,
  sd: 1.0,
  sp: 1.1,
}
const jolly = {
  name: 'ようき',
  at: 1.0,
  df: 1.0,
  sa: 0.9,
  sd: 1.0,
  sp: 1.1,
}
const naive = {
  name: 'むじゃき',
  at: 1.0,
  df: 1.0,
  sa: 1.0,
  sd: 0.9,
  sp: 1.1,
}

export const nature = {
  namespaced: true,
  state: {
    current: bashful,
    natures: [bashful,
      lonely, adamant, naughty, brave,
      bold, impish, lax, relaxed,
      modest, mild, rash, quiet,
      calm, gentle, careful, sassy,
      timid, hasty, jolly, naive]
  },
  getters: {
    currentName: state => state.current.name,
    attack: (state) => state.current.at,
    defense: (state) => state.current.df,
    spAttack: (state) => state.current.sa,
    spDefense: (state) => state.current.sd,
    speed: (state) => state.current.sp,
  },
  mutations: {
    setNature (state, natureName) {
      let next = state.natures.find((v) => v.name == natureName);
      state.current = next;
    }
  }
}

