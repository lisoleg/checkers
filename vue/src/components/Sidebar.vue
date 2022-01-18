<template>
  <SpSidebar v-on:sidebar-open="sidebarOpen = true" v-on:sidebar-close="sidebarOpen = false">
    <template v-slot:default>
      <SpLinkIcon link="/" text="控制面板" icon="Dashboard" />
      <SpLinkIcon link="/relayers" text="区块浏览器" icon="Transactions" />
    </template>
    <template v-slot:footer>
      <SpButton type="primary" @click="changeLanguage('ch')">中   文</SpButton>
      <SpButton type="primary" @click="changeLanguage('en')">English</SpButton>
       <h3>@{{$t('menus.metab')}}</h3>
    </template>
  </SpSidebar>
</template>

<script>
import { getCurrentInstance } from "vue";
import { useI18n } from 'vue-i18n'//要在js中使用国际化
export default {
  name: 'Sidebar',
  data() {
    return {
      sidebarOpen: true,
    }
  },
  computed: {
    hasWallet() {
      return this.$store.hasModule(['common', 'wallet'])
    },
  },

  setup() {
    const { proxy } = getCurrentInstance();
    function changeLanguage(type) {
      proxy.$i18n.locale = type;
    }
    const { t } = useI18n();
    return { changeLanguage };
  },
}
</script>
