<template>
  <!-- <div class="page-heading about-heading">
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          
        </div>
      </div>
    </div>
  </div> -->

  <div class="products">
    <div class="container">
      <div class="text-content buffer-top-text">
        <!-- <h4>Thông tin chi tiết</h4> -->
        <h2>Thông tin chi tiết Sản Phẩm</h2>
      </div>
      <div class="row buffer-top">
        <div class="col-md-4 col-xs-12">
          <div>
            <img
              :src="`http://localhost:3000/${product.Image}`"
              alt=""
              class="img-fluid wc-image"
            />
          </div>
        </div>

        <div class="col-md-8 col-xs-12">
          <div   class="form">
            <h2>{{ product.Name }}</h2>
            <br />
            <p class="lead">
              <!-- <small><del> $999.00</del></small><strong class="text-primary">$779.00</strong> -->
              <strong class="text-primary">{{ money }} </strong>
            </p>

            <br />

            <br />

            <div class="row">
              <div class="col-sm-4">
                <label class="control-label">Loại sản phẩm</label>
                <div class="form-group">
                  <select class="form-control">
                    <option value="0">18 gears</option>
                    <option value="1">21 gears</option>
                    <option value="2">27 gears</option>
                  </select>
                </div>
              </div>
              <div class="col-sm-8">
                <label class="control-label">Số lượng</label>

                <div class="row">
                  <div class="col-sm-6">
                    <div class="form-group">
                      <input type="text" class="form-control" placeholder="1"  v-model.number="quantityOfProduct"/>
                    </div>
                  </div>

                  <div class="col-sm-6">
                    <button
                     
                      class="btn btn-primary btn-block"
                      @click="addToCart(product)"
                      >Add to Cart</button
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <h2>Mô tả sản phẩm</h2>
        <p class="lead">
          {{ product.Description }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductDetail",
  props: ["id"],
  data() {
    return {
      product: {},
      quantityOfProduct:1,
    };
  },

  async created() {
    console.log(this.id);
    const response = await fetch(
      `http://localhost:3000/api/product/${this.id}`
    );
    this.product = await response.json();
    console.log(this.product);
  },

  methods: {
    addToCart(product) {
      // this.quantityOfProduct=event.target.value;
      product["quantity"] = this.quantityOfProduct;
      this.$store.commit("addProductToCart", product);
    },
  },

  computed: {
    money() {
      const formatter = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      });
      return formatter.format(this.product.Price);
    },

    
  },
};
</script>

<style scoped>
.page-heading {
  padding: 210px 0px 90px 0px;
  text-align: center;
  background-position: center center;
  background-size: cover;
  position: relative;
  background-repeat: no-repeat;
  background-image: url("../assets/images/background/heading-1-1920x500.jpg");
}

.page-heading:before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.7);
}

/* .page-heading .text-content {
  position: relative;
} */

/* .page-heading .text-content h4 {
  color: #f33f3f;
  font-size: 22px;
  text-transform: uppercase;
  font-weight: 700;
  margin-bottom: 15px;
} */
/* .page-heading .text-content h4 {
  color: #fff;
  font-size: 62px;
  text-transform: uppercase;
  letter-spacing: 5px;
} */

.buffer-top-text {
  margin-top: 100px;
  border-bottom: 1px solid rgb(163, 163, 163);
}
.buffer-top {
  margin-top: 50px;
}
.row {
  margin-bottom: 20px;
}
</style>
