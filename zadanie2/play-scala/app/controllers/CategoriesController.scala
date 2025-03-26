package controllers

import javax.inject._
import play.api._
import play.api.libs.json._
import play.api.mvc._
import scala.util.Random
import play.api.libs.json.Reads._
import play.api.libs.functional.syntax._

case class Category(id:Int, name:String)

object Category{

  implicit val categoryWrites : Writes[Category] = new Writes[Category] {
    def writes(category: Category) = Json.obj(
      "id"  -> category.id,
      "name" -> category.name,
    )
  }

  implicit val categoryReads: Reads[Category] = new Reads[Category] {
      def reads(json: JsValue): JsResult[Category] = {
        val id = (json \ "id").asOpt[Int].getOrElse(Random.nextInt(1000)) 
        val name = (json \ "name").asOpt[String]

        (name) match {
          case Some(n) =>
            JsSuccess(Category(id, n))
          case _ =>
            JsError("Invalid category data")
        }
      }
    }
}
@Singleton
class CategoriesController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

private var categories = List(
    Category(0, "Komputery i laptopy"),
    Category(1, "Akcesoria komputerowe"),
    Category(2, "Sprzet AGD"),
    Category(3, "sluchawki"),
  )

  def getAllCategories: Action[AnyContent] = Action {
    val json = Json.toJson(categories)
    Ok(json)
  }

  def getCategory(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound("Category not found")
    }
  }

  def deleteCategory(id:Int): Action[AnyContent] = Action {
    if (categories.exists(_.id == id))
    {
      categories = categories.filterNot(_.id == id)
      Ok(Json.toJson(categories))
    }
    else
    {
      NotFound("Category not found")
    }
  }

  def updateCategory: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category] match {
      case JsSuccess(newCategory, _) =>
        if(!categories.exists(_.id == newCategory.id)){
          categories = categories :+ newCategory
          Created(Json.toJson(categories)) 
        }else
        {
          categories = categories.map {
            case category if category.id == newCategory.id => category.copy(name = newCategory.name)
            case category => category
          }
          Ok(Json.toJson(categories))
        }
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid category data", "errors" -> errorMessages))
    }
  }


  def addCategory: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category] match {
      case JsSuccess(newCategory, _) =>
        categories = categories :+ newCategory
        Created(Json.toJson(newCategory)) 
      case JsError(errors) =>
        val errorMessages = errors.map {
        case (path, validationErrors) =>
          (path.toString(), validationErrors.map(_.message).mkString(", "))
      }
        BadRequest(Json.obj("message" -> "Invalid category data", "errors" -> errorMessages))
    }
  }

}
