<template>
  <section>
    <template v-for="provider in carouselData">
      <provider-card class="generic-card" :provider="provider" :key="provider.id" />
    </template>
  </section>
</template>

<script>
import ProviderCard from "./ProviderCard.vue";
export default {
  components: { ProviderCard },
  props: {
    carouselData: {
      type: Array,
      required: true
    }
  },
  data() {
    return {
      providers: [
        {
          id: "603a33feeae9a825587c142b",
          contact: {
            contact_person: "",
            email: "lol@lmfao.com",
            number: "+919234234324",
            address: {
              street_address_1: "Sai Sumukha Sameeksha Apartments",
              street_address_2: "Rachenahalli Main Road",
              city: "Bengaluru",
              state: "Karnataka",
              country: "",
              landmark: "",
              zip_code: "560077",
              coordinates: [13.0559579, 77.63091969999999],
            },
            other: [],
            "website": "https://gaurav.app"
          },
          name: "Couselling service BLR",
          service_type: "Counselling",
          fee_range: "less than Rs 500",
          fee_negotiable: "Yes/No/Maybe",
          timings: "24x7",
          online: false
        },
        {
            "id": "603a366e278a3e72d4c0f406",
            "contact": {
                "contact_person": "",
                "email": "lol@lmfao.com",
                "number": "+919234234324",
                "address": {
                    "street_address_1": "Sai Sumukha Sameeksha Apartments",
                    "street_address_2": "Rachenahalli Main Road",
                    "city": "Bengaluru",
                    "state": "Karnataka",
                    "country": "",
                    "landmark": "",
                    "zip_code": "560077"
                },
                "other": [],
                "website": "https://gaurav.app"
            },
            "name": "My online service",
            "service_type": "couns",
            "fee_range": "less than Rs 500",
            "fee_negotiable": "Yes/No/Maybe",
            "timings": "24x7",
            "online": true
        },
      ],
    };
  },
};
</script>