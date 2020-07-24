import scala.concurrent.Future
import scala.concurrent.ExecutionContext.Implicits.global

class Message(var id: Int, var title: String)

class Trial {
  val m1 = new Message(1, "satu")
  val m2 = new Message(2, "dua")
  val m3 = new Message(3, "tiga")
  val mAllowed = List(1,2)
  val ms = getList()
  val diffFms = List()
  val diffMs = ms.map(_.id).diff(mAllowed)

  def trial() = {

    println(s"result > ${m1.title}")
    println(s"id > ${ms.map(_.id).mkString(",")}")
    println(s"title > ${ms.map(_.title).mkString(",")}")
    println(s"wrap single > ${wrapString("test")}")
    println(s"wrap bulk > ${ms.map( m => (wrapString(m.title), m.id.toString))}")
    println(s"ms diffs > ${diffMs.mkString(",")}")
    for {
      dfms <- getFutureList() map { messages =>
        val diffFms = messages.map(_.id).diff(mAllowed)
        if (diffFms.length != 0) {
          println(s"id ${diffFms.mkString(",")} not allowed")
        }
        diffFms
      }
    } yield {
      println(s"diff fms title > ${dfms.mkString(",")}")
    }

    for {
      rfms <- getFutureList()
      ffms <- getFutureListBySeq(rfms)
    } yield {
      println(s"forwarded fms title > ${ffms.map(_.title).mkString(",")}")
    }

    for {
      rfms <- getFutureList()
      ffms <- getFutureListByPointer(rfms:_*)
    } yield {
      println(s"forwarded pointer fms title > ${ffms.map(_.title).mkString(",")}")
    }

    val testMap = Map("test-1" -> 1, "test-2" -> 2)
    val (testMapStr1, _)= testMap.foldLeft(("[",1)){ case((s, idx), (str, num)) =>
      val suffix = if (idx == testMap.size) "]" else ","
      (s"$s($str,$num)$suffix", idx+1)
    }
    val testMapStr2 = testMap map { case(key, value) => s"($key, $value)"} mkString(",")
    val testMapStr3 = testMap.map{case(key, value) => (key, value)}(collection.breakOut).mkString(",")
    val testMapStr4 = testMap.map{case(key, value) => (key, value)}(collection.breakOut) mkString(",")
    println(s"test-map-1 > ${testMapStr1}")
    println(s"test-map-2 > ${testMapStr2}")
    println(s"test-map-3 > ${testMapStr3}")
    println(s"test-map-4 > ${testMapStr4}")


    val testMapTuple = Map("test-1" -> (true, None), "test-2" -> (false, Some(2)), "test-3" -> (false, Some(3)))
    val (successItems, failedItems) = testMapTuple.foldLeft(Map.empty[String, Boolean], Map.empty[String, Int]) {
      case ((successItem, failedItem), (id, (ok, amount))) =>
        if (ok) {
          (successItem + (id -> ok), failedItem)
        } else {
          (successItem, failedItem + (id -> amount.get))
        }
    }
    println(s"test-map-tuple ok > ${successItems}")
    println(s"test-map-tuple !ok > ${failedItems}")

    val regionCode = "34.02.05.0001"
    val codes = regionCode.split("\\.")
    val districtRegionCode = codes.take(3).mkString(".")
    println(s"district > $districtRegionCode")

    val minValue = List(10, 91).min
    println(s"min-value : $minValue")
  }


  def getFutureList(): Future[Seq[Message]] = Future {
    Seq(m1, m2, m3)
  }

  def getFutureListBySeq(raw:Seq[Message]): Future[Seq[Message]] = Future {
    raw
  }

  def getFutureListByPointer(raw:Message*): Future[Seq[Message]] = Future {
    raw
  }

  def getList(): Seq[Message] = Seq(m1, m2, m3)


  def wrapString(raw:String): String = Seq("p",raw).mkString("-")
}
