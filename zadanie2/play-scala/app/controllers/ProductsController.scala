package controllers

import javax.inject._
import play.api._
import play.api.libs.json._
import play.api.mvc._
import scala.util.Random
import play.api.libs.json.Reads._
import play.api.libs.functional.syntax._

case class Product(id:Int, name:String, price : Double)

object Product{

  implicit val ProductWrites: Writes[Product] = new Writes[Product] {
    def writes(product: Product) = Json.obj(
      "id"  -> product.id,
      "name" -> product.name,
      "price" -> product.price
    )
  }

  implicit val ProductReads: Reads[Product] = new Reads[Product] {
      def reads(json: JsValue): JsResult[Product] = {
        val id = (json \ "id").asOpt[Int].getOrElse(Random.nextInt(1000)) 
        val name = (json \ "name").asOpt[String]
        val price = (json \ "price").asOpt[Double]

        ( name, price) match {
          case (Some(n), Some(p)) =>
            JsSuccess(Product(id, n, p))
          case _ =>
            JsError("Invalid product data")
        }
      }
    }
}

@Singleton
class ProductsController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  private var products = List(
    Product(1, "Laptop", 2500.0),
    Product(2, "Myszka", 50.0)
  )

  def index() = Action { implicit request: Request[AnyContent] =>
    Ok(views.html.index())
  }

  def getAllProducts: Action[AnyContent] = Action {
    val json = Json.toJson(products)

    Ok(json)
  }

  def getProductById(id: Int): Action[AnyContent] = Action {
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound("Product not found")
    }
  }

  def deleteProductById(id:Int): Action[AnyContent] = Action {
    if (products.exists(_.id == id))
    {
      products = products.filterNot(_.id == id)
      Ok(Json.toJson(products))
    }
    else
    {
      NotFound("Product not found")
    }
  }

  def updateProduct: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product] match {
      case JsSuccess(newProduct, _) =>
        if(!products.exists(_.id == newProduct.id)){
          products = products :+ newProduct
          Created(Json.toJson(newProduct)) 
        }else
        {
          products = products.map {
            case product if product.id == newProduct.id => product.copy(name = newProduct.name, price = newProduct.price)
            case product => product
          }
          Ok(Json.toJson(products))
        }
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid product data", "errors" -> errorMessages))
    }
  }


  def addProduct: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product] match {
      case JsSuccess(newProduct, _) =>
        products = products :+ newProduct
        Created(Json.toJson(newProduct)) 
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid product data", "errors" -> errorMessages))
    }
  }

}
