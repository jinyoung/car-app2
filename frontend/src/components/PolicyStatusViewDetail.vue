<template>

    <v-card outlined>
        <v-card-title>
            PolicyStatus # {{$route.params.id }}
        </v-card-title>

        <v-card-text>
            <div>
                <String label="Status" v-model="item.status" :editMode="editMode" @change="change" />
            </div>
        </v-card-text>
    </v-card>

</template>

<script>
  const axios = require('axios').default;

  export default {
    name: 'PolicyStatusViewDetail',
    props: {
      value: Object,
    },
    data: () => ({
        item : [],
    }),
    async created() {
      var params = this.$route.params;
      var temp = await axios.get(axios.fixUrl('/policyStatuses/' + params.id))

      this.item = temp.data;

    },
    methods: {
    }
  }
</script>

