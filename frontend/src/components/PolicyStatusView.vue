<template>

    <v-data-table
        :headers="headers"
        :items="policyStatus"
        :items-per-page="5"
        class="elevation-1"
    ></v-data-table>

</template>

<script>
    const axios = require('axios').default;

    export default {
        name: 'PolicyStatusView',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
                { text: "id", value: "id" },
                { text: "status", value: "status" },
            ],
            policyStatus : [],
        }),
          async created() {
            var temp = await axios.get(axios.fixUrl('/policyStatuses'))

            temp.data._embedded.policyStatuses.map(obj => obj.id=obj._links.self.href.split("/")[obj._links.self.href.split("/").length - 1])

            this.policyStatus = temp.data._embedded.policyStatuses;
        },
        methods: {
        }
    }
</script>

