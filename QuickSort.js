function quickSort(array, left = 0, right = array.length - 1) {
    // console.log("~~~~~~~~~~~~~~~~~~~~~~",left,right,"11111111111111111111");

    if (left >= right) {
        // console.log("~~~~~~~~~~~~~~~~~~~~~~",left,right);
      return;
    }
    const mid = Math.floor((left + right) / 2);
    const pivot = array[mid];
    // console.log("~~~##########################",left,right);

    const partition = divide(array, left, right, pivot);
    // console.log(array,"11@122121212121")
    quickSort(array, left, partition - 1);
    // console.log(array, partition, right,"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",left)
    quickSort(array, partition, right);
    // console.log(array, partition, right,"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!",left,"@222")
  
    function divide(array, left, right, pivot) {
//         console.log("array:",array ,"left:",left, "right:",right, "pivot:",pivot);
        while (left <= right) {
            while (array[left] < pivot) {
            left++;
            }
            while (array[right] > pivot) {
            right--;
            }
//             console.log("left:",left, "right:",right);
//             console.log("------before","array:left",array[left],"array:right",array[right],array);
            if (left <= right) {
            let temp = array[left];
            array[left] = array[right];
            array[right] = temp;
            left++;
            right--;
            }
//             console.log("------after","array:left",array[left],"array:right",array[right] ,array);

        }   
      return left;
    }
    return array;
}
  console.log(quickSort([9, 12,2, 1,5, 26, 7, 14, 3, 7,20]));


  @Entity
public class User {
    @Id
    private Long id;
    private String name;
    
    protected User() {}
    
    public User(Long id, String name) {
        this.id = id;
        this.name = name;
    }
    
    public Long getId() {
        return this.id;
    }
    
    public String getName() {
        return this.name;
    }
}
public interface UserRepository {
    void save(User user);
    User findById(long id);
}
public class FakeUserRepository implements UserRepository {
    private Collection<User> users = new ArrayList<>();
    
    @Override
    public void save(User user) {
        if (findById(user.getId()) == null) {
            user.add(user);
        }
    }
    
    @Override
    public User findById(long id) {
        for (User user : users) {
            if (user.getId() == id) {
                return user;
            }
        }
        return null;
    }
}