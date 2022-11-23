
import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);


import PolicyApplicationManager from "./components/listers/PolicyApplicationCards"
import PolicyApplicationDetail from "./components/listers/PolicyApplicationDetail"

import PolicyStatusView from "./components/PolicyStatusView"
import PolicyStatusViewDetail from "./components/PolicyStatusViewDetail"
import PolicyHistoryManager from "./components/listers/PolicyHistoryCards"
import PolicyHistoryDetail from "./components/listers/PolicyHistoryDetail"


export default new Router({
    // mode: 'history',
    base: process.env.BASE_URL,
    routes: [
            {
                path: '/policyApplications',
                name: 'PolicyApplicationManager',
                component: PolicyApplicationManager
            },
            {
                path: '/policyApplications/:id',
                name: 'PolicyApplicationDetail',
                component: PolicyApplicationDetail
            },

            {
                path: '/policyStatuses',
                name: 'PolicyStatusView',
                component: PolicyStatusView
            },
            {
                path: '/policyStatuses/:id',
                name: 'PolicyStatusViewDetail',
                component: PolicyStatusViewDetail
            },
            {
                path: '/policyHistories',
                name: 'PolicyHistoryManager',
                component: PolicyHistoryManager
            },
            {
                path: '/policyHistories/:id',
                name: 'PolicyHistoryDetail',
                component: PolicyHistoryDetail
            },



    ]
})
