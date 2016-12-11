import java.io.Serializable;
import java.util.ArrayList;

public class InvertedIndex implements Serializable {
	//Inverted Index holds a Document ID paired with documentRef
	//string -> List[documentRef]

	//object documentRef {
    //	name: string
    //	path: string
    //	occurrences: List[FileLocation]
	//}

	private String documentID;
	private ArrayList<documentRef> documents;
    
    private InvertedIndex() {

    }

    private class documentRef {
    	//object documentRef {
	    //	name: string
	    //	path: string
	    //	occurrences: List[FileLocation]
		//}
		private String name;
		private String path;
		private ArrayList<Integer> occurrences;

		public void setName(String name) {
			this.name = name;
		}

		public void setPath(String path) {
			this.path = path;
		}

		public void addOccurrence(int loc) {
			this.occurrences.add(loc);
		}

		public String getName() {
			return name;
		}

		public String getPath() {
			return path;
		}


    }
}
