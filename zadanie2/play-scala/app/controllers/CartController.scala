package controllers

import javax.inject._
import play.api._
import play.api.libs.json._
import play.api.mvc._
import play.api.libs.json.Reads._
import play.api.libs.functional.syntax._

case class CartItem(productId:Int, quantity : Int)

object CartItem{

  implicit val CartItemWrites: Writes[CartItem] = new Writes[CartItem] {
    def writes(cartItem: CartItem) = Json.obj(
      "productId"  -> cartItem.productId,
      "quantity" -> cartItem.quantity,
    )
  }

  implicit val cartItemReads: Reads[CartItem] = Json.reads[CartItem]
}

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  private var cart = List(
    CartItem(1,2),
    CartItem(2,4)
  )
  
  def getAllCartItems: Action[AnyContent] = Action {
    val json = Json.toJson(cart)

    Ok(json)
  }

  def getCartItem(id: Int): Action[AnyContent] = Action {
    cart.find(_.productId == id) match {
      case Some(cartItem) => Ok(Json.toJson(cartItem))
      case None => NotFound("cartItem not found")
    }
  }

  def deleteCartItem(id:Int): Action[AnyContent] = Action {
    if (cart.exists(_.productId == id))
    {
      cart = cart.filterNot(_.productId == id)
      Ok(Json.toJson(cart))
    }
    else
    {
      NotFound("cartItem not found")
    }
  }

  def updateCart: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem] match {
      case JsSuccess(newCartItem, _) =>
        if(!cart.exists(_.productId == newCartItem.productId)){
          cart = cart :+ newCartItem
          Created(Json.toJson(cart)) 
        }else
        {
          cart = cart.map {
            case cartItem if cartItem.productId == newCartItem.productId => cartItem.copy(quantity = newCartItem.quantity )
            case cartItem => cartItem
          }
          Ok(Json.toJson(cart))
        }
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid cart item data", "errors" -> errorMessages))
    }
  }


  def addCartItem: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem] match {
      case JsSuccess(newCartItem, _) =>
        cart = cart :+ newCartItem
        Created(Json.toJson(newCartItem)) 
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid cartItem data", "errors" -> errorMessages))
    }
  }

}
